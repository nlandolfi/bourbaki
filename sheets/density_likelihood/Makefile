# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make density_likelihood.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

density_likelihood.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > density_likelihood.tex

density_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex density_likelihood.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode density_likelihood.tex
	../../latexrun density_likelihood.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: density_likelihood.pdf
	open density_likelihood.pdf

o:
	make open

terms: density_likelihood.tex
	bbk_sheet -mode ts density_likelihood.tex

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
	rm -f density_likelihood.*
	make reset
	bbk_sheet -mode mk > Makefile