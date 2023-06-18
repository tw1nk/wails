(()=>{var fe=Object.defineProperty;var g=(e,t)=>{for(var n in t)fe(e,n,{get:t[n],enumerable:!0})};var S={};g(S,{SetText:()=>pe,Text:()=>we});var de=window.location.origin+"/wails/runtime";function me(e,t,n){let i=new URL(de);i.searchParams.append("method",e),n&&i.searchParams.append("args",JSON.stringify(n));let o={headers:{}};return t&&(o.headers["x-wails-window-name"]=t),new Promise((l,u)=>{fetch(i,o).then(a=>{if(a.ok)return a.headers.get("Content-Type")&&a.headers.get("Content-Type").indexOf("application/json")!==-1?a.json():a.text();u(Error(a.statusText))}).then(a=>l(a)).catch(a=>u(a))})}function r(e,t){return function(n,i=null){return me(e+"."+n,t,i)}}var N=r("clipboard");function pe(e){N("SetText",{text:e})}function we(){return N("Text")}var W={};g(W,{Hide:()=>ge,Quit:()=>xe,Show:()=>he});var E=r("application");function ge(){E("Hide")}function he(){E("Show")}function xe(){E("Quit")}var k={};g(k,{Log:()=>Ce});var ve=r("log");function Ce(e){return ve("Log",e)}var O={};g(O,{GetAll:()=>be,GetCurrent:()=>Se,GetPrimary:()=>Me});var L=r("screens");function be(){return L("GetAll")}function Me(){return L("GetPrimary")}function Se(){return L("GetCurrent")}var Ee="useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict";var h=(e=21)=>{let t="",n=e;for(;n--;)t+=Ee[Math.random()*64|0];return t};var We=r("call"),f=new Map;function ke(){let e;do e=h();while(f.has(e));return e}function F(e,t,n){let i=f.get(e);i&&(n?i.resolve(JSON.parse(t)):i.resolve(t),f.delete(e))}function B(e,t){let n=f.get(e);n&&(n.reject(t),f.delete(e))}function G(e,t){return new Promise((n,i)=>{let o=ke();t=t||{},t["call-id"]=o,f.set(o,{resolve:n,reject:i}),We(e,t).catch(l=>{i(l),f.delete(o)})})}function I(e){return G("Call",e)}function U(e,t,...n){return G("Call",{packageName:"wails-plugins",structName:e,methodName:t,args:n})}function Y(e){let t=r("window",e);return{Center:()=>void t("Center"),SetTitle:n=>void t("SetTitle",{title:n}),Fullscreen:()=>void t("Fullscreen"),UnFullscreen:()=>void t("UnFullscreen"),SetSize:(n,i)=>t("SetSize",{width:n,height:i}),Size:()=>t("Size"),SetMaxSize:(n,i)=>void t("SetMaxSize",{width:n,height:i}),SetMinSize:(n,i)=>void t("SetMinSize",{width:n,height:i}),SetAlwaysOnTop:n=>void t("SetAlwaysOnTop",{alwaysOnTop:n}),SetPosition:(n,i)=>t("SetPosition",{x:n,y:i}),Position:()=>t("Position"),Screen:()=>t("Screen"),Hide:()=>void t("Hide"),Maximise:()=>void t("Maximise"),Show:()=>void t("Show"),Close:()=>void t("Close"),ToggleMaximise:()=>void t("ToggleMaximise"),UnMaximise:()=>void t("UnMaximise"),Minimise:()=>void t("Minimise"),UnMinimise:()=>void t("UnMinimise"),Restore:()=>void t("Restore"),SetBackgroundColour:(n,i,o,l)=>void t("SetBackgroundColour",{r:n,g:i,b:o,a:l})}}var Le=r("events"),z=class{constructor(t,n,i){this.eventName=t,this.maxCallbacks=i||-1,this.Callback=o=>(n(o),this.maxCallbacks===-1?!1:(this.maxCallbacks-=1,this.maxCallbacks===0))}},x=class{constructor(t,n=null){this.name=t,this.data=n}},s=new Map;function v(e,t,n){let i=s.get(e)||[],o=new z(e,t,n);return i.push(o),s.set(e,i),()=>Oe(o)}function Q(e,t){return v(e,t,-1)}function X(e,t){return v(e,t,1)}function Oe(e){let t=e.eventName,n=s.get(t).filter(i=>i!==e);n.length===0?s.delete(t):s.set(t,n)}function j(e){console.log("dispatching event: ",{event:e});let t=s.get(e.name);if(t){let n=[];t.forEach(i=>{i.Callback(e)&&n.push(i)}),n.length>0&&(t=t.filter(i=>!n.includes(i)),t.length===0?s.delete(e.name):s.set(e.name,t))}}function J(e,...t){[e,...t].forEach(i=>{s.delete(i)})}function q(){s.clear()}function C(e){Le("Emit",e)}var ze=r("dialog"),d=new Map;function Re(){let e;do e=h();while(d.has(e));return e}function V(e,t,n){let i=d.get(e);i&&(n?i.resolve(JSON.parse(t)):i.resolve(t),d.delete(e))}function _(e,t){let n=d.get(e);n&&(n.reject(t),d.delete(e))}function m(e,t){return new Promise((n,i)=>{let o=Re();t=t||{},t["dialog-id"]=o,d.set(o,{resolve:n,reject:i}),ze(e,t).catch(l=>{i(l),d.delete(o)})})}function K(e){return m("Info",e)}function Z(e){return m("Warning",e)}function $(e){return m("Error",e)}function p(e){return m("Question",e)}function ee(e){return m("OpenFile",e)}function te(e){return m("SaveFile",e)}var ye=r("contextmenu");function Te(e,t,n,i){return ye("OpenContextMenu",{id:e,x:t,y:n,data:i})}function ie(e){e?window.addEventListener("contextmenu",ne):window.removeEventListener("contextmenu",ne)}function ne(e){oe(e.target,e)}function oe(e,t){let n=e.getAttribute("data-contextmenu");if(n)t.preventDefault(),Te(n,t.clientX,t.clientY,e.getAttribute("data-contextmenu-data"));else{let i=e.parentElement;i&&oe(i,t)}}function re(e,t=null){let n=new x(e,t);C(n)}function Ae(){document.querySelectorAll("[data-wml-event]").forEach(function(t){let n=t.getAttribute("data-wml-event"),i=t.getAttribute("data-wml-confirm"),o=t.getAttribute("data-wml-trigger")||"click",l=function(){if(i){p({Title:"Confirm",Message:i,Buttons:[{Label:"Yes"},{Label:"No",IsDefault:!0}]}).then(function(u){u!=="No"&&re(n)});return}re(n)};t.removeEventListener(o,l),t.addEventListener(o,l)})}function le(e){wails.Window[e]===void 0&&console.log("Window method "+e+" not found"),wails.Window[e]()}function De(){document.querySelectorAll("[data-wml-window]").forEach(function(t){let n=t.getAttribute("data-wml-window"),i=t.getAttribute("data-wml-confirm"),o=t.getAttribute("data-wml-trigger")||"click",l=function(){if(i){p({Title:"Confirm",Message:i,Buttons:[{Label:"Yes"},{Label:"No",IsDefault:!0}]}).then(function(u){u!=="No"&&le(n)});return}le(n)};t.removeEventListener(o,l),t.addEventListener(o,l)})}function R(){Ae(),De()}var y=function(e){chrome.webview.postMessage(e)};var ae=new Map;function se(e){let t=new Map;for(let[n,i]of Object.entries(e))typeof i=="object"&&i!==null?t.set(n,se(i)):t.set(n,i);return t}fetch("/wails/flags").then(e=>{e.json().then(t=>{ae=se(t)})});function Pe(e){let t=e.split("."),n=ae;for(let i of t)if(n instanceof Map?n=n.get(i):n=n[i],n===void 0)break;return n}function b(e){return Pe(e)}var w=!1;function He(e){let t=window.getComputedStyle(e.target).getPropertyValue("--webkit-app-region");return t&&(t=t.trim()),t!=="drag"||e.buttons!==1?!1:e.detail===1}function ue(){window.addEventListener("mousedown",Fe),window.addEventListener("mousemove",Ge),window.addEventListener("mouseup",Be)}var M=null;function Ne(e){return M?(y("resize:"+M),!0):!1}function Fe(e){if(!Ne())if(He(e)){if(e.offsetX>e.target.clientWidth||e.offsetY>e.target.clientHeight)return;w=!0}else w=!1}function Be(e){(e.buttons!==void 0?e.buttons:e.which)>0&&T()}function T(){document.body.style.cursor="default",w=!1}function c(e){document.documentElement.style.cursor=e||Ie,M=e}function Ge(e){if(w){w=!1,(e.buttons!==void 0?e.buttons:e.which)>0&&y("drag");return}Ue(e)}var Ie="auto";function Ue(e){let t=b("system.resizeHandleHeight")||5,n=b("system.resizeHandleWidth")||5,i=b("resizeCornerExtra")||3,o=window.outerWidth-e.clientX<n,l=e.clientX<n,u=e.clientY<t,a=window.outerHeight-e.clientY<t,A=window.outerWidth-e.clientX<n+i,D=e.clientX<n+i,P=e.clientY<t+i,H=window.outerHeight-e.clientY<t+i;!l&&!o&&!u&&!a&&M!==void 0?c():A&&H?c("se-resize"):D&&H?c("sw-resize"):D&&P?c("nw-resize"):P&&A?c("ne-resize"):l?c("w-resize"):u?c("n-resize"):a?c("s-resize"):o&&c("e-resize")}window.wails={...ce(null),Capabilities:{}};fetch("/wails/capabilities").then(e=>{e.json().then(t=>{window.wails.Capabilities=t})});window._wails={dialogCallback:V,dialogErrorCallback:_,dispatchWailsEvent:j,callCallback:F,callErrorCallback:B,endDrag:T};function ce(e){return{Clipboard:{...S},Application:{...W,GetWindowByName(t){return ce(t)}},Log:k,Screens:O,Call:I,Plugin:U,WML:{Reload:R},Dialog:{Info:K,Warning:Z,Error:$,Question:p,OpenFile:ee,SaveFile:te},Events:{Emit:C,On:Q,Once:X,OnMultiple:v,Off:J,OffAll:q},Window:Y(e)}}console.log("Wails v3.0.0 Debug Mode Enabled");ie(!0);ue();document.addEventListener("DOMContentLoaded",function(e){R()});})();
