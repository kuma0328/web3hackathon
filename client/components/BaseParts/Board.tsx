import { ReactNode } from 'react';
type TProps = {
  children: ReactNode;
  className?: string;
};
export const Board = ({ children, className }: TProps) => {
  return (
    <div
      className={`${className} relative m-10 rounded-md border-2 border-black bg-white shadow-lg`}
    >
      <span className="absolute left-1/3  -my-3 h-7 w-1/3 rounded-md bg-amber-600"></span>
      {children}
    </div>
  );
};
