import React, { ReactNode } from 'react';
type TProps = {
  children: ReactNode;
  className?: string;
};
export const Descriptions = ({ children, className }: TProps) => {
  return (
    <div
      className={`${className} my-5 border-t-2 border-b-2 border-black py-5 text-left`}
    >
      {children}
    </div>
  );
};
