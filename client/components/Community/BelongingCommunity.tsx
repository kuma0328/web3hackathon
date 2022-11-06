import Image from 'next/image';
type TProps = {
  className?: string;
};

export const BelongingCommunity = ({ className }: TProps) => {
  return (
    <div className="flex gap-10 overflow-x-auto">
      {[1, 2].map((t) => {
        return (
          <Image
            src={`/images/mock/${t}.jpg`}
            width={100}
            height={100}
            alt={''}
            className=" my-3 rounded-md border-2 border-black"
          />
        );
      })}
    </div>
  );
};
