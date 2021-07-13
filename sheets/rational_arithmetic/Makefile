# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rational_arithmetic.pdf

rational_arithmetic.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rational_arithmetic.tex

rational_arithmetic.pdf: ../../*.tex ../../trademark.pdf *.tex rational_arithmetic.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rational_arithmetic.tex
	../../latexrun rational_arithmetic.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rational_arithmetic.pdf
	open rational_arithmetic.pdf

o:
	make open

terms: rational_arithmetic.tex
	bbk_sheet -mode ts rational_arithmetic.tex

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
	rm -f rational_arithmetic.*
	make reset
	bbk_sheet -mode mk > Makefile