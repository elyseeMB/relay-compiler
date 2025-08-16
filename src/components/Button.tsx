type ParamsProps = {
  title: string
}

export function Button({title}: ParamsProps) {
  return <button>
    {title}
  </button>;
}