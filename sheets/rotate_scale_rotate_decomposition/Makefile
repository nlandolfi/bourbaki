# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rotate_scale_rotate_decomposition.pdf

rotate_scale_rotate_decomposition.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rotate_scale_rotate_decomposition.tex

rotate_scale_rotate_decomposition.pdf: ../../*.tex ../../trademark.pdf *.tex rotate_scale_rotate_decomposition.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rotate_scale_rotate_decomposition.tex
	../../latexrun rotate_scale_rotate_decomposition.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rotate_scale_rotate_decomposition.pdf
	open rotate_scale_rotate_decomposition.pdf

o:
	make open

terms: rotate_scale_rotate_decomposition.tex
	bbk_sheet -mode ts rotate_scale_rotate_decomposition.tex

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
	rm -f rotate_scale_rotate_decomposition.*
	make reset
	bbk_sheet -mode mk > Makefile