import { useCommunityStore } from '../../stores/CommunityStore';

export const CommunityInputs = () => {
  const { image, title, description, setImage, setTitle, setDescription } =
    useCommunityStore();
  return (
    <div className="p-5">
      <div className="  my-2 w-full rounded-md border py-32  text-center">
        <input
          type="file"
          value={image}
          onChange={(e) => setImage(e.target.value)}
        />
      </div>

      <input
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        className=" w-full rounded-md border border-black p-1"
        placeholder="コミュニティ名か料理名を記入してください"
      />

      <input
        type="text"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        className=" my-5 w-full rounded-md border border-black px-1 pb-20 pt-1"
        placeholder="コミュニティや料理の説明を記入してください"
      />
    </div>
  );
};
