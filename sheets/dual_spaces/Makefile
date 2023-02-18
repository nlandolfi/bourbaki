# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dual_spaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

dual_spaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dual_spaces.tex

dual_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex dual_spaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dual_spaces.tex
	../../latexrun dual_spaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dual_spaces.pdf
	open dual_spaces.pdf

o:
	make open

terms: dual_spaces.tex
	bbk_sheet -mode ts dual_spaces.tex

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
	rm -f dual_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile