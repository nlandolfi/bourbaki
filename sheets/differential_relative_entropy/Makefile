# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make differential_relative_entropy.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

differential_relative_entropy.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > differential_relative_entropy.tex

differential_relative_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex differential_relative_entropy.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode differential_relative_entropy.tex
	../../latexrun differential_relative_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_relative_entropy.pdf
	open differential_relative_entropy.pdf

o:
	make open

terms: differential_relative_entropy.tex
	bbk_sheet -mode ts differential_relative_entropy.tex

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
	rm -f differential_relative_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile