# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_balls.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_balls.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_balls.tex

real_balls.pdf: ../../*.tex ../../trademark.pdf *.tex real_balls.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_balls.tex
	../../latexrun real_balls.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_balls.pdf
	open real_balls.pdf

o:
	make open

terms: real_balls.tex
	bbk_sheet -mode ts real_balls.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_balls.*
	make reset
	bbk_sheet -mode mk > Makefile