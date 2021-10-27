# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probably_approximately_correct_learners.pdf

probably_approximately_correct_learners.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probably_approximately_correct_learners.tex

probably_approximately_correct_learners.pdf: ../../*.tex ../../trademark.pdf *.tex probably_approximately_correct_learners.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probably_approximately_correct_learners.tex
	../../latexrun probably_approximately_correct_learners.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probably_approximately_correct_learners.pdf
	open probably_approximately_correct_learners.pdf

o:
	make open

terms: probably_approximately_correct_learners.tex
	bbk_sheet -mode ts probably_approximately_correct_learners.tex

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
	rm -f probably_approximately_correct_learners.*
	make reset
	bbk_sheet -mode mk > Makefile