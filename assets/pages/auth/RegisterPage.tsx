import React from "preact/compat";
import { useToast } from "../../hooks/useToast.tsx";
import { Field } from "../../components/ui/Field.tsx";
import { buildEndpoint } from "../../providers/RelayProviders.tsx";
import { useFormWithSchema } from "../../hooks/useFormWithSchema.ts";

import { z } from "zod";
import { useNavigate } from "react-router";

const schema = z.object({
  fullname: z.string(),
  password: z.string(),
});

export default function RegisterPage() {
  const navigate = useNavigate();
  const { pushToast } = useToast();

  const { register, formState, handleSubmit } = useFormWithSchema(schema, {
    defaultValues: {
      fullname: "",
      password: "",
    },
  });

  const onSubmit = handleSubmit(async (data) => {
    const response = await fetch(
      buildEndpoint("/api/console/v1/auth/register"),
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(data),
      }
    );

    if (response.ok) {
      console.log("ok");
    }

    navigate("/", { replace: true });
  });

  return (
    <div className="stack">
      <div className="is-flex is-justify-content-center is-align-items-center">
        <form onSubmit={onSubmit as any}>
          <h2 className="pb-4 is-size-5 color">Registration</h2>

          <Field
            htmlFor="fullName"
            placeholder="your fullname"
            label="Fullname"
            {...register("fullname")}
            error={formState.errors.fullname?.message}
          />

          <Field
            htmlFor="password"
            placeholder="your password"
            label="Password"
            {...register("password")}
            error={formState.errors.password?.message}
          />

          <div className="field">
            <div className="control">
              <button className="button is-link">Submit</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
}
