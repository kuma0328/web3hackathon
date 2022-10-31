import { ReactNode } from 'react';
type TProps = {
  chatTitle: string;
  children: ReactNode;
  className?: string;
};
export const Chat = ({ children, className, chatTitle }: TProps) => {
  return (
    <div
      className={`absolute w-full rounded-md border-2 border-black bg-white shadow-lg`}
    >
      <span
        className={`${className} absolute z-30 -my-3 mx-auto flex h-10 w-1/3 items-center justify-center rounded-md transition-all hover:z-50 hover:-my-7 hover:h-14`}
        onClick={() => console.log('click')}
      >
        {chatTitle}
      </span>
      {children}
    </div>
  );
};
