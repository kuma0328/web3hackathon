import { Button } from '../BaseParts/Button';
import { Input } from '../BaseParts/Input';
type TProps = {
  mail: string;
  password: string;
  name?: string;
  passwordConfirmation?: string;
  onChangeMail: () => void;
  onChangePassword: () => void;
  onChangeName?: () => void;
  onChangePasswordConfirmation?: () => void;
  onSubmit: () => void;
};
export const Form = ({
  mail,
  password,
  passwordConfirmation,
  name,
  onChangePassword,
  onChangeMail,
  onChangeName,
  onChangePasswordConfirmation,
  onSubmit,
}: TProps) => {
  return (
    <div className="w-1/3">
      <Input
        id="name"
        name="名前"
        type="text"
        value={name}
        onChange={onChangeName}
      />
      <Input
        id="mail"
        name="メールアドレス"
        type="mail"
        value={mail}
        onChange={onChangeMail}
      />
      <Input
        id="password"
        name="パスワード"
        type="password"
        value={password}
        onChange={onChangePassword}
      />
      <Input
        id="password-confirmation"
        name="パスワード確認"
        type="password"
        value={passwordConfirmation}
        onChange={onChangePasswordConfirmation}
      />
      <div className="my-5 flex justify-center">
        <Button onClick={onSubmit}>ユーザー登録</Button>
      </div>
    </div>
  );
};
