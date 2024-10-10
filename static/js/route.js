function logout() {
  document.cookie = null
  window.location.replace("/login")
}
function go_back() {
  window.location.replace("/homepage")
}
function open_msg() {
  window.location.replace("/msg_center")
}
function open_ai_chat() {
  window.location.replace("/ai_chat_room")
}

let currentTheme = localStorage.getItem('theme') || 0;
const themes = [
  {
    body_bg_color: "#f7f2e0", /* 暖色背景 */
    text_color: "#333",/* 文字颜色 */
    container_bg_color: "#282c34", /* 科技感背景 */
    sidebar_bg_color: "#383e4b", /* 侧边栏背景颜色 */
    sidebar_heading_color: "#ffbe76",/* 侧边栏标题颜色 */
    header_bg_color: "#ffbe76",/* 头部背景颜色 */
    notification_bg_color: "#444f61",/* 通知中心背景颜色 */
    notification_item_bg_color: "#383e4b", /* 通知项背景颜色 */
    notification_item_system_text_color: "#ff6b6b",/* 通知项system文字颜色 */
    notification_item_user_text_color: "#ffffff", /* 通知项user文字颜色 */
    button_bg_color: "#ffbe76", /* 按钮背景颜色 */
    footer_bg_color: "#383e4b", /* 底部背景颜色 */
    footer_text_color: "#fff", /* 底部文本颜色 */
    table_header_bg_color: '#383e4b', /* 表头背景颜色 */
    table_cell_bg_color: "#444f61", /* 单元格背景颜色 */
    closed_color: "#28a745",/* 绿色表示已关闭 */
    open_color: "#ff6b6b", /* 红色表示未关闭 */
    notification_list_bg_color:"#2c313a",
  },
  {
    body_bg_color: " #ffffff", 
    text_color: "#f4fbff",
    container_bg_color: " #f7f7fc",
    sidebar_bg_color: "#343dec",
    sidebar_heading_color: "#f3f3ff",
    header_bg_color: "#006dff",
    notification_bg_color: "#ffffff",
    notification_item_bg_color: "#ffffff",
    notification_item_system_text_color: "#ff6b6b",
    notification_item_user_text_color: "#000000",
    button_bg_color: "#006dff",
    footer_bg_color: "#343dec",
    footer_text_color: "#ffffff",
    table_header_bg_color: '#383e4b',
    table_cell_bg_color: "#444f61",
    closed_color: "#28a745",
    open_color: "#ff6b6b",
    notification_list_bg_color:"#f7f7fc",
  }
];
function applyTheme(theme) {
  document.documentElement.style.setProperty('--body-bg-color', theme.body_bg_color);
  document.documentElement.style.setProperty('--text-color', theme.text_color);
  document.documentElement.style.setProperty('--container-bg-color', theme.container_bg_color);
  document.documentElement.style.setProperty('--sidebar-bg-color', theme.sidebar_bg_color);
  document.documentElement.style.setProperty('--sidebar-heading-color', theme.sidebar_heading_color);
  document.documentElement.style.setProperty('--header-bg-color', theme.header_bg_color);
  document.documentElement.style.setProperty('--notification-bg-color', theme.notification_bg_color);
  document.documentElement.style.setProperty('--notification-item-bg-color', theme.notification_item_bg_color);
  document.documentElement.style.setProperty('--notification-item-system-text-color', theme.notification_item_system_text_color);
  document.documentElement.style.setProperty('--notification-item-user-text-color', theme.notification_item_user_text_color);
  document.documentElement.style.setProperty('--button-bg-color', theme.button_bg_color);
  document.documentElement.style.setProperty('--footer-bg-color', theme.footer_bg_color);
  document.documentElement.style.setProperty('--footer-text-color', theme.footer_text_color);
  document.documentElement.style.setProperty('--table-header-bg-color', theme.table_header_bg_color);
  document.documentElement.style.setProperty('--table-cell-bg-color', theme.table_cell_bg_color);
  document.documentElement.style.setProperty('--closed-color', theme.closed_color);
  document.documentElement.style.setProperty('--open-color', theme.open_color);
  document.documentElement.style.setProperty('--notification-list-bg-color', theme.notification_list_bg_color);
}
function switchTheme() {
  console.log("switchTheme")
  currentTheme = (currentTheme + 1) % themes.length;
  localStorage.setItem('theme', currentTheme);
  applyTheme(themes[currentTheme]);
}
var logo = document.getElementById("logo_switchTheme")
logo.addEventListener("click", switchTheme)
document.addEventListener("DOMContentLoaded", applyTheme(themes[currentTheme]))