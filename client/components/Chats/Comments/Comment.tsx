import Image from 'next/image';

export const Comment = () => {
  return (
    <div className="my-5 flex px-5">
      <Image
        src="/images/mock/user.png"
        width={50}
        height={50}
        alt={''}
        className="m-auto mb-5 rounded-full"
      />
      <div className="mx-3 flex-1 rounded-md bg-gray-300 p-5">
        <p>レシピ共有しました！！</p>
      </div>
    </div>
  );
};
