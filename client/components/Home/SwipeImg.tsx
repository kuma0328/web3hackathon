import Image from 'next/image';
type TProps = {
  image: number;
};
export const SwipeImg = ({ image }: TProps) => {
  return (
    <div className="h-full rounded-full bg-white object-cover p-2 md:p-5 ">
      <Image
        src={`/images/${image}.jpg`}
        width={900}
        height={900}
        alt="cuisine"
        className="rounded-full drop-shadow-2xl"
      />
    </div>
  );
};
