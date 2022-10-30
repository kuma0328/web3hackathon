import Link from 'next/link';
import { ReactNode } from 'react';
type TProps = {
  link: string;
  className?: string;
  children: ReactNode;
};
export const LinkTo = ({ link, className, children }: TProps) => {
  return (
    <div
      className={`${className} rounded-md border-2 border-black py-2 px-4 font-medium shadow-md transition hover:scale-110 hover:transform hover:bg-black hover:text-white`}
    >
      <Link href={link}>{children}</Link>
    </div>
  );
};
