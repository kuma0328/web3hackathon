import Link from 'next/link';
import { ReactNode } from 'react';
type TProps = {
  link: string;
  className?: string;
  children: ReactNode;
};
export const LinkTo = ({ link, className, children }: TProps) => {
  console.log(link);
  return (
    <Link
      passHref
      href={{
        pathname: link,
        query: { communityId: '1' },
      }}
      className={`${className} rounded-md border-2 border-black py-2 px-4 font-medium shadow-md transition hover:scale-110 hover:transform hover:bg-black hover:text-white`}
    >
      {children}
    </Link>
  );
};
