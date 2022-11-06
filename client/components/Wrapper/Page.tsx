import { ReactNode } from 'react';
import { useGetUser } from '../../hooks/User/useGetUser';

type TProps = {
  children: ReactNode;
  className?: string;
  wide: boolean;
};

export const Page = ({ className, children, wide }: TProps) => {
  const { data, error, isLoading } = useGetUser(1);
  return (
    <div
      className={
        wide
          ? `${className} flex min-h-screen flex-col items-center justify-center bg-bc p-3`
          : `${className} bg-bc`
      }
    >
      {children}
    </div>
  );
};
