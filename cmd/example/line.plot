set term svg size width,height enhanced mouse font 'Verdana,14'
set output filetitle

set title headtitle

set datafile separator ","

set xdata time

set timefmt '%Y-%m-%d'


set format x "%d-%m\n%Y"

set grid


set xzeroaxis linetype 3 linewidth 1.5


set style line 1 linetype 17 linecolor rgb "black" linewidth 1.000

set samples 1000


plot filename using 1:2 with lines ls 1 title "My trend"

     
