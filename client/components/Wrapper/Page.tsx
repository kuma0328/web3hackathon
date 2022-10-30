import { ReactNode } from 'react';

type TProps = {
  children: ReactNode;
  className?: string;
};
export const Page = ({ className, children }: TProps) => {
  return (
    <div
      className={`${className} flex min-h-screen flex-col items-center justify-center bg-bc p-3`}
    >
      {children}
    </div>
  );
};
