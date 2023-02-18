# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_square_roots.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_square_roots.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_square_roots.tex

real_square_roots.pdf: ../../*.tex ../../trademark.pdf *.tex real_square_roots.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_square_roots.tex
	../../latexrun real_square_roots.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_square_roots.pdf
	open real_square_roots.pdf

o:
	make open

terms: real_square_roots.tex
	bbk_sheet -mode ts real_square_roots.tex

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
	rm -f real_square_roots.*
	make reset
	bbk_sheet -mode mk > Makefile