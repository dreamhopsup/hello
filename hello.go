Msat  = 1000e3
Aex   = 10e-12

// define exchange length
lex := sqrt(10e-12 / (0.5 * mu0 * pow(1000e3 ,2)))

d     := 30 * lex                        // we test for d/lex = 30
Sizex := 5*d                             // magnet size x
Sizey := 1*d
Sizez := 0.1*d

nx := pow(2, ilogb(Sizex / (0.75*lex)))  // power-of-two number of cells
ny := pow(2, ilogb(Sizey / (0.75*lex)))  // not larger than 0.75 exchange lengths

SetGridSize(nx, ny, 1)
SetCellSize(Sizex/nx, Sizey/ny, Sizez)

m = Uniform(1, 0.1, 0)                   // initial mag
relax() 

save(m)                                  // remanent magnetization
print("<m> for d/lex=30: ", m.average())
