type TProps = {
  id?: string;
  name?: string;
  type?: string;
  value?: string;
  onChange?: () => void;
};
export const Input = ({ id, name, type, value, onChange }: TProps) => {
  return (
    <div className="mx-5 grid grid-cols-1 justify-items-start py-3">
      <label htmlFor={id} className="text-left">
        {name}
      </label>

      <input
        id={id}
        type={type}
        onChange={onChange}
        value={value}
        className="w-full rounded-lg border-2 border-black"
      />
    </div>
  );
};
