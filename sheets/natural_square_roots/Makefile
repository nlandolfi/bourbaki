# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_square_roots.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

natural_square_roots.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > natural_square_roots.tex

natural_square_roots.pdf: ../../*.tex ../../trademark.pdf *.tex natural_square_roots.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_square_roots.tex
	../../latexrun natural_square_roots.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_square_roots.pdf
	open natural_square_roots.pdf

o:
	make open

terms: natural_square_roots.tex
	bbk_sheet -mode ts natural_square_roots.tex

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
	rm -f natural_square_roots.*
	make reset
	bbk_sheet -mode mk > Makefile