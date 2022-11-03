import { ReactNode } from 'react';
type TProps = {
  children: ReactNode;
  className?: string;
  onClick?: () => void;
};
export const Button = ({ children, className, onClick }: TProps) => {
  return (
    <button
      className={`${className} rounded-md border-2 border-black py-2 px-4 font-medium shadow-md transition hover:scale-110 hover:transform hover:bg-black hover:text-white`}
      onClick={onClick}
    >
      {children}
    </button>
  );
};
