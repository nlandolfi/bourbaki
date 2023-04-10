# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make convergence_in_measure.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

convergence_in_measure.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > convergence_in_measure.tex

convergence_in_measure.pdf: ../../*.tex ../../trademark.pdf *.tex convergence_in_measure.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode convergence_in_measure.tex
	../../latexrun convergence_in_measure.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convergence_in_measure.pdf
	open convergence_in_measure.pdf

o:
	make open

terms: convergence_in_measure.tex
	bbk_sheet -mode ts convergence_in_measure.tex

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
	rm -f convergence_in_measure.*
	make reset
	bbk_sheet -mode mk > Makefile