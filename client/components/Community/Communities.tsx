import { Community } from './Community';

export const Communities = () => {
  return (
    <div className=" grid grid-cols-5">
      {[0, 1, 2, 3, 4, 5].map(() => {
        return <Community />;
      })}
    </div>
  );
};
