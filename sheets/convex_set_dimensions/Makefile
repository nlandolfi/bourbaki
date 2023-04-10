# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make convex_set_dimensions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

convex_set_dimensions.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > convex_set_dimensions.tex

convex_set_dimensions.pdf: ../../*.tex ../../trademark.pdf *.tex convex_set_dimensions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode convex_set_dimensions.tex
	../../latexrun convex_set_dimensions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convex_set_dimensions.pdf
	open convex_set_dimensions.pdf

o:
	make open

terms: convex_set_dimensions.tex
	bbk_sheet -mode ts convex_set_dimensions.tex

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
	rm -f convex_set_dimensions.*
	make reset
	bbk_sheet -mode mk > Makefile