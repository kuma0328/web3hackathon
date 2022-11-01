import { ReactNode } from 'react';
type TProps = {
  chatTitle: string;
  children: ReactNode;
  className?: string;
  onClick: () => void;
};
export const Chat = ({ children, className, chatTitle, onClick }: TProps) => {
  return (
    <div className="absolute w-full rounded-md bg-white shadow-lg">
      <button
        className={`${className} absolute z-30 -my-3 mx-auto flex h-10 w-1/3 items-center justify-center rounded-md transition-all hover:z-50 hover:-my-7 hover:h-14`}
        onClick={onClick}
      >
        {chatTitle}
      </button>
      {children}
    </div>
  );
};
