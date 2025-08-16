type Params = {
  title: string;
  description: string;
}

export function Message({title, description}: Params) {
  return <div>
    <h4>{title}</h4>
    <p>
      {description}
    </p>
  </div>;
}