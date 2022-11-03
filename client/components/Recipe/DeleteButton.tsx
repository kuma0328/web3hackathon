type TProps = {
  data: string;
  onDelete: (data: string) => void;
};
export const DeleteButton = ({ data, onDelete }: TProps) => {
  return (
    <button
      className="absolute -top-2 -right-2 h-6 w-6 rounded-md bg-orange-400"
      onClick={() => onDelete(data)}
    >
      ×
    </button>
  );
};
