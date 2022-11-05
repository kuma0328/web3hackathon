type TProps = {
  id?: string;
  name?: string;
  type?: string;
  value?: string;
  onChange: (e: string) => void;
};
export const Input = ({ id, name, type, value, onChange }: TProps) => {
  console.log(value);
  return (
    <div className="py-3">
      <label htmlFor={id} className="text-left">
        {name}
      </label>

      <input
        id={id}
        type={type}
        onChange={(e) => onChange(e.target.value)}
        value={value}
        className="w-full rounded-lg border-2 border-black p-1"
      />
    </div>
  );
};
