import React, { forwardRef } from "preact/compat";
import { InputHTMLAttributes } from "preact/compat";
import { $ZodCheckGreaterThan } from "zod/v4/core";

type Params = {
  htmlFor: string;
  label: string;
  error?: string;
} & InputHTMLAttributes<HTMLInputElement>;

export const Field = forwardRef<HTMLInputElement, Params>(
  (
    {
      label,
      htmlFor,
      name,
      type = "text",
      placeholder = "Enter your value",
      error,
      ...props
    },
    ref
  ) => {
    const defaultHtmlFor = !htmlFor ? name : htmlFor;

    return (
      <div className="field">
        <label htmlFor={defaultHtmlFor} className="label">
          {label}
        </label>
        <div className="control">
          <input
            ref={ref}
            className="input"
            name={name}
            type={type}
            placeholder={placeholder}
            {...props}
          />
          <span>{error}</span>
        </div>
      </div>
    );
  }
);
