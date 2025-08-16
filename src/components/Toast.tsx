type ParamsProps = {
  title: string;
  content: string;
};

export function Toast({ title, content }: ParamsProps) {
  return (
    <div className="notification is-info">
      <div className="is-flex-direction-column">
        <span>
          <strong>{title}</strong>
        </span>
        {content}
      </div>
    </div>
  );
}
