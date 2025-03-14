from customtkinter import *


class Skeleton(CTk):
    def __init__(self, fg_color=None, **kwargs):
        super().__init__(fg_color, **kwargs)
        CTkCheckBox(self).pack()
        CTkEntry(self).pack()
        CTkEntry(self).pack()
        CTkButton(self).pack()
        CTkSlider(self).pack()
        self.mainloop()


Skeleton()
