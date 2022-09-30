# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_arithmetic.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_arithmetic.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_arithmetic.tex

real_arithmetic.pdf: ../../*.tex ../../trademark.pdf *.tex real_arithmetic.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_arithmetic.tex
	../../latexrun real_arithmetic.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_arithmetic.pdf
	open real_arithmetic.pdf

o:
	make open

terms: real_arithmetic.tex
	bbk_sheet -mode ts real_arithmetic.tex

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
	rm -f real_arithmetic.*
	make reset
	bbk_sheet -mode mk > Makefile