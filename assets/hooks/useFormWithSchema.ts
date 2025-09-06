import { useForm, UseFormProps } from "react-hook-form";
import type { z, ZodType } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

export function useFormWithSchema<T extends ZodType<any, any, any>>(
  schema: T,
  options?: Omit<UseFormProps<z.infer<T>>, "resolver">
) {
  return useForm<z.infer<T>>({
    ...options,
    resolver: zodResolver(schema),
  });
}
