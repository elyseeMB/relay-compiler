import { InputHTMLAttributes } from "react";

type Params = {
  htmlFor: string;
  label: string;
  name: string;
  placeholder: string;
  type?: InputHTMLAttributes<HTMLInputElement>["type"];
};

export function Field({
  label,
  htmlFor,
  name,
  type = "text",
  placeholder = "Enter your value",
}: Params) {
  const defaultHtmlFor = !htmlFor ? name : htmlFor;
  return (
    <>
      <div className="field">
        <label htmlFor={defaultHtmlFor} className="label">
          {label}
        </label>
        <div className="control">
          <input
            className="input"
            name={name}
            type={type}
            placeholder={placeholder}
          />
        </div>
      </div>
    </>
  );
}
