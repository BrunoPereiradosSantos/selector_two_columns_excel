package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"

	"gui_golang_react_excel/handlers"

	"github.com/xuri/excelize/v2"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

const (
	SKU      = 0
	QUANTITY = 16
)

type FileStatus struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (a *App) ProcessExcel(fileData []byte) (string, error) {
	file, err := excelize.OpenReader(bytes.NewReader(fileData))
	if err != nil {
		fs := FileStatus{Message: "Erro ao processar o arquivo", Status: http.StatusInternalServerError}
		json, _ := handlers.ConvertToJson(fs)
		return json, errors.New("Erro ao processar o Excel: " + err.Error())
	}

	sheetName := file.GetSheetName(0)

	if sheetName == "" {
		fs := FileStatus{Message: "Nenhuma planilha encontrada no arquivo", Status: http.StatusBadRequest}
		json, _ := handlers.ConvertToJson(fs)
		return json, errors.New("Nenhuma planilha encontrada no arquivo")
	}

	rows, err := file.GetRows(sheetName)

	newFile := excelize.NewFile()
	newSheet := "Sheet1"
	newFile.SetSheetName("Sheet1", newSheet)
	newFile.SetCellValue(newSheet, "A1", "SKU")
	newFile.SetCellValue(newSheet, "B1", "QTD")
	slicedRows := rows[2:]

	for i, row := range slicedRows {
		newFile.SetCellValue(newSheet, fmt.Sprintf("A%d", i+2), row[SKU])
		newFile.SetCellValue(newSheet, fmt.Sprintf("B%d", i+2), row[QUANTITY])
	}

	var buf bytes.Buffer

	if err := newFile.Write(&buf); err != nil {
		fs := FileStatus{Message: "Erro ao criar o novo arquivo excelvo novo arquivo excel", Status: http.StatusInternalServerError}
		json, _ := handlers.ConvertToJson(fs)

		return json, errors.New(err.Error())
	}

	outputPath := "sku_quantity.xlsx"

	path, err := handlers.SaveFileToDownloads(outputPath, buf.Bytes())
	if err != nil {
		fs := FileStatus{Message: "Erro ao salvar o arquivo ", Status: http.StatusInternalServerError}
		json, _ := handlers.ConvertToJson(fs)
		return json, nil
	}

	fs := FileStatus{Message: "Arquivo Salvo em: " + path, Status: http.StatusCreated}

	json, _ := handlers.ConvertToJson(fs)

	return json, nil
}
