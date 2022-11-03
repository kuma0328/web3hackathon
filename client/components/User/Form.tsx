import { Button } from '../BaseParts/Button';
import { Input } from '../BaseParts/Input';
import { TypoGraphy } from '../BaseParts/TypoGraphy';
type TProps = {
  title?: string;
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
  title,
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
    <div className="m-10">
      <TypoGraphy className="text-center text-xl">{title}</TypoGraphy>
      {name !== undefined && (
        <Input
          id="name"
          name="名前"
          type="text"
          value={name}
          onChange={onChangeName}
        />
      )}
      <Input
        id="mail"
        name="メールアドレス"
        type="email"
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
      {passwordConfirmation !== undefined && (
        <Input
          id="password-confirmation"
          name="パスワード確認"
          type="password"
          value={passwordConfirmation}
          onChange={onChangePasswordConfirmation}
        />
      )}

      <div className="my-5 flex justify-center">
        <Button onClick={onSubmit}>{title}する</Button>
      </div>
    </div>
  );
};
