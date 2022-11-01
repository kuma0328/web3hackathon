export const CommunityInputs = () => {
  return (
    <div className="p-5">
      <div className="  my-2 w-full rounded-md border py-32  text-center">
        <input type="file" name="" id="" />
      </div>
      <input
        type="text"
        name=""
        id=""
        className=" w-full rounded-md border border-black p-1"
        placeholder="コミュニティ名か料理名を記入してください"
      />

      <input
        type="text"
        className=" my-5 w-full rounded-md border border-black px-1 pb-20 pt-1"
        placeholder="コミュニティや料理の説明を記入してください"
      />
    </div>
  );
};
