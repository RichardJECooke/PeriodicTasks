export function formatDate(date: Date): string {
  const weekdays = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
  const weekday = weekdays[date.getDay()];
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${weekday} ${year}-${month}-${day}`;
}

export function getDayFromDate(date: Date): string {
  return formatDate(date).split(' ')[0];
}

export function getDateFromDate(date: Date): string {
  return formatDate(date).split(' ')[1];
}
