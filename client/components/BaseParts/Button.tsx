import { ReactNode } from 'react';
type TProps = {
  children: ReactNode;
  onClick?: () => void;
};
export const Button = ({ children, onClick }: TProps) => {
  return (
    <button
      className="rounded-md border-2 border-black py-3 px-5 font-medium shadow-md transition hover:scale-110 hover:transform hover:bg-black hover:text-white"
      onClick={onClick}
    >
      {children}
    </button>
  );
};
