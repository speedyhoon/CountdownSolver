# CountdownSolver
Solve math problems presented in the TV game shows [Countdown](https://en.wikipedia.org/wiki/Countdown_(game_show)#Numbers_round), Countdown Masters, [8 Out of 10 Cats Does Countdown](https://en.wikipedia.org/wiki/8_Out_of_10_Cats_Does_Countdown), Letters and Numbers and [Des chiffres et des lettres](https://en.wikipedia.org/wiki/Des_chiffres_et_des_lettres).

Rules:
---
* Each number can only be used once,
* Addition, subtraction and multiplication can be used,
* Division can be used where the result is a whole number.

Issues:
---
* Brakets have not been implemented.  
  Given the numbers: **100**, **50**, **4**, **1**, **3**, **10** and the target: **932**,  
  The program incorrectly ouputs: ```100 - 4 - 3 × 10 + 1 = 931```  
  Instead of: ```100 * 10 - ((50 + 1) * 4 / 3) = 932```
