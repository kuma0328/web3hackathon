import { ReactNode } from 'react';
type TProps = {
  children: ReactNode;
  className?: string;
};
export const TypoGraphy = ({ className, children }: TProps) => {
  return (
    <div className={`${className}  font-semibold text-fc`}>{children}</div>
  );
};
