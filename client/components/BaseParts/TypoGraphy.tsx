import { ReactNode } from 'react';
type TProps = {
  children: ReactNode;
  size?: string;
};
export const TypoGraphy = ({ size, children }: TProps) => {
  return <div className={`${size} font-semibold text-fc`}>{children}</div>;
};
