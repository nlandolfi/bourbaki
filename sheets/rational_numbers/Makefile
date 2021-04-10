# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make rational_numbers.pdf
	make graph.pdf

rational_numbers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rational_numbers.tex

rational_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex rational_numbers.tex
	# pdflatex --file-line-error -interaction=nonstopmode rational_numbers.tex
	../../latexrun rational_numbers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rational_numbers.pdf
	open rational_numbers.pdf

o:
	make open

terms: rational_numbers.tex
	bbk_sheet -mode ts rational_numbers.tex

spell:
	aspell -c rational_numbers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f rational_numbers.*
	make reset
	bbk_sheet -mode mk > Makefile