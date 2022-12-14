import { ReactNode } from 'react';
import { Header } from '../ Header/ Header';

type TProps = {
  children: ReactNode;
  className?: string;
  wide: boolean;
};

export const Page = ({ className, children, wide }: TProps) => {
  return (
    <div
      className={
        wide
          ? `${className} flex min-h-screen flex-col items-center justify-center bg-bc p-3`
          : `${className} bg-bc`
      }
    >
      <Header />
      {children}
    </div>
  );
};
