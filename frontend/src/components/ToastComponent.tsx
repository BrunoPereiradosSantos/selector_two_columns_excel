import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export const ToastSuccess = (message: string) => {
  toast.success(message, {
    autoClose: 750,
  });
};

export const toastError = (message: string) => {
  toast.error(message, {
    autoClose: 750,
  });
};

const ToastComponent = () => {
  return (
    <section>
      <ToastContainer position="top-center" />
    </section>
  );
};

export default ToastComponent;
