import Link from 'next/link';
import Image from 'next/image';
export const Header = () => {
  return (
    <div className="absolute top-0 left-0">
      <Link href="./community">
        <Image
          src={`/images/Logo.png`}
          width={100}
          height={100}
          alt="cuisine"
        />
      </Link>
    </div>
  );
};
