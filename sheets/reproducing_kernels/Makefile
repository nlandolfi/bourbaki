# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make reproducing_kernels.pdf

reproducing_kernels.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > reproducing_kernels.tex

reproducing_kernels.pdf: ../../*.tex ../../trademark.pdf *.tex reproducing_kernels.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode reproducing_kernels.tex
	../../latexrun reproducing_kernels.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: reproducing_kernels.pdf
	open reproducing_kernels.pdf

o:
	make open

terms: reproducing_kernels.tex
	bbk_sheet -mode ts reproducing_kernels.tex

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
	rm -f reproducing_kernels.*
	make reset
	bbk_sheet -mode mk > Makefile