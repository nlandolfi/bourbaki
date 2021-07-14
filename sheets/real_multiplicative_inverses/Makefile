# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_multiplicative_inverses.pdf

real_multiplicative_inverses.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_multiplicative_inverses.tex

real_multiplicative_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex real_multiplicative_inverses.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_multiplicative_inverses.tex
	../../latexrun real_multiplicative_inverses.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_multiplicative_inverses.pdf
	open real_multiplicative_inverses.pdf

o:
	make open

terms: real_multiplicative_inverses.tex
	bbk_sheet -mode ts real_multiplicative_inverses.tex

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
	rm -f real_multiplicative_inverses.*
	make reset
	bbk_sheet -mode mk > Makefile