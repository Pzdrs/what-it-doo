function formatDateOrTime(input: Date | string | number): string {
  const date = input instanceof Date ? input : new Date(input);
  const now = new Date();

  const isToday =
    date.getFullYear() === now.getFullYear() &&
    date.getMonth() === now.getMonth() &&
    date.getDate() === now.getDate();

  if (isToday) {
    // Return only the time, e.g. "14:35"
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
  } else {
    // Return only the date, e.g. "2025-09-26"
    return date.toLocaleDateString([], {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    });
  }
}

export { formatDateOrTime };