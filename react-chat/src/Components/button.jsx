const InteractionButton = ({ isLike, count, onAction }) => {
  // Define colors based on the button type
  const activeColor = isLike 
    ? "bg-blue-600 hover:bg-blue-700" 
    : "bg-gray-600 hover:bg-gray-700";

  return (
    <button
      onClick={() => onAction(isLike)}
      className={`${activeColor} text-white font-semibold py-3 px-6 rounded-lg shadow-md transition-all active:scale-95 flex items-center gap-2`}
    >
      <span>{isLike ? "👍" : "👎"}</span>
      <span>{isLike ? "Like" : "Dislike"}</span>
      <span className="bg-white/20 px-2 py-0.5 rounded text-sm ml-1">
        {count}
      </span>
    </button>
  );

};

export default InteractionButton

