import savana from './assets/images/savana.svg';
import dots from './assets/images/dots.svg';
import './App.css';

import { ProcessExcel } from '../wailsjs/go/main/App';
import ToastComponent, {
  toastError,
  ToastSuccess,
} from './components/ToastComponent';
import { useState } from 'react';

type ResponseType = {
  message: string;
  status: number;
};

function App() {
  const [loading, setIsLoading] = useState(false);

  const processForm = async (
    e: React.FormEvent<HTMLFormElement>,
  ): Promise<void> => {
    e.preventDefault();

    setIsLoading(true);

    const supportedTypes = [
      'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
    ];

    const fileInput = e.currentTarget.elements.namedItem(
      'file',
    ) as HTMLInputElement;

    if (!fileInput || !fileInput.files || fileInput.files.length === 0) {
      toastError('Selecione um Arquivo');
      setIsLoading(false);
      return;
    }

    if (!supportedTypes.includes(fileInput.files[0].type)) {
      toastError('Formato de arquivo não suportado');
      setIsLoading(false);
      return;
    }

    const file = fileInput.files[0];

    const arrayBuffer = await file.arrayBuffer();

    try {
      const response = await ProcessExcel(
        Array.from(new Uint8Array(arrayBuffer!)),
      );

      const toJson: ResponseType = JSON.parse(response);

      if (toJson.status !== 201) {
        throw new Error(toJson.message);
      }

      if (toJson.status === 201) {
        ToastSuccess(toJson.message);
      }
    } catch (error) {
      toastError(error as string);
      setIsLoading(false);
      return;
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div id="App">
      <ToastComponent />
      <img src={savana} id="logo" alt="logo" />
      {!loading && (
        <form onSubmit={processForm} encType="multipart/form-data">
          <input type="file" name="file" accept=".xlsx" />

          <button className="btn" type="submit">
            Processar Excel
          </button>
        </form>
      )}
      {loading && <img src={dots} alt="dots from loading" />}
    </div>
  );
}

export default App;
