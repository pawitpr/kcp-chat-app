# Simple Chat App
This is a simple chat app use kcp protocol and kcp is a  KCP is a fast and reliable ARQ (Automatic Repeat-reQuest) protocol that uses a different automatic retransmission policy than TCP and has a lower network latency than TCP. In practice, KCP over UDP is often used instead of TCP for online games and audio/video transmission.
# Demo video


https://github.com/pawitpr/kcp-chat-app/assets/123424956/f52e2b0d-9d6e-4346-9c69-753fba775fb3

https://github.com/pawitpr/kcp-chat-app/assets/123424956/f2eb1137-97fd-4318-a72c-fa014519a1ac





# Installation
You can download the executable file from here 
| Version             | Link                                                                |
| ----------------- | ------------------------------------------------------------------ |
| Window-x64 | [64bit version](https://firebasestorage.googleapis.com/v0/b/cloudstorage-pawit.appspot.com/o/kcp-chat-app-windowsx64.zip?alt=media&token=08141ff5-5d46-4a64-8c95-f4e5051cb8e5)  |
| Window-x86 | [32 bit version](https://firebasestorage.googleapis.com/v0/b/cloudstorage-pawit.appspot.com/o/kcp-chat-app-windowsx86.zip?alt=media&token=1e8b6ffb-3b1c-4bcc-aa1e-dabdca43eba5)  |
| Linux-Arm-64 | [64 bit arm version](https://firebasestorage.googleapis.com/v0/b/cloudstorage-pawit.appspot.com/o/kcp-chat-app-linux-arm-64.zip?alt=media&token=4f7b75ef-bd83-44ff-ae9a-f12a84bb7a8d)  |
| Linux-x86_64 | [64 bit and 86 bit for intel or amd processor](https://firebasestorage.googleapis.com/v0/b/cloudstorage-pawit.appspot.com/o/kcp-chat-app-linux-x86_64.zip?alt=media&token=3741dc12-12b7-4b90-bc5d-94f506784197)  |

please check your arch and download the correct file
for checking arch in linux:
```bash
 uname -m 
 # if it show aarch64 then your processor arch is arm64 or if show x86_64 then your processor arch is x86_64
 ```
 
Installation for linux

```bash
  chmod +x server 
  chmod +x client
 ```
 Installtion log: 
 [buildlog.txt](https://github.com/pawitpr/kcp-chat-app/files/11526172/buildlog.txt)
 
