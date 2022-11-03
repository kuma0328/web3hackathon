type TProps = {
  data: string;
  onCreate: (data: string) => void;
  resetCurrent: () => void;
};
export const CreateButton = ({ data, onCreate, resetCurrent }: TProps) => {
  return (
    <button
      className="ml-2 rounded-md border border-black p-2"
      onClick={() => {
        onCreate(data);
        resetCurrent();
      }}
    >
      ＋
    </button>
  );
};
