# chattingServer
채팅 서버를 만들기 위해서 tcp 소켓 통신 부터 차근차근 만들어보자

## TCP 소켓 통신
1. client가 주기적으로 서버로 'hey server'라는 message를 날린다. 
2. server에서는 받은 메세지를 'server got :'을 덧씌워서 client에게 다시 보내준다. 

## simple chat service
1. 방에 들어가기 (join)
2. 닉네임 설정하기 (nickname)
3. 방에서 모든 사람에게 채팅하기 (msg)
4. 방에서 나가기 (quit)

![Screen-Recording-2022-12-11-at-3 02 51-PM (1)](https://user-images.githubusercontent.com/64356885/206889960-5a3e936e-04b0-4b30-98ac-089793cb883e.gif)
